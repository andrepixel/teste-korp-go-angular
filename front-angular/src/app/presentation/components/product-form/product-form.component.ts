import { Component, OnInit } from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  FormBuilder,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { ActivatedRoute, Router } from '@angular/router';
import { ProductService } from '../../../core/services/product.service';
import { Product } from '../../../core/models/product.model';

@Component({
  selector: 'app-product-form',
  standalone: true,
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatInputModule,
    MatButtonModule,
    MatCardModule,
  ],
  templateUrl: './product-form.component.html',
  styleUrls: ['./product-form.component.scss'],
})
export class ProductFormComponent implements OnInit {
  productForm!: FormGroup;
  isEditMode = false;
  productId: string | null = null;

  constructor(
    private fb: FormBuilder,
    private productService: ProductService,
    private route: ActivatedRoute,
    private router: Router
  ) {}

  ngOnInit() {
    this.productForm = this.fb.group({
      name: ['', Validators.required],
      price: [1, [Validators.required, Validators.min(1)]],
      available: [
        1,
        [
          Validators.required,
          Validators.min(1),
          Validators.pattern('^[0-9]*$'),
        ],
      ],
    });

    this.productId = this.route.snapshot.paramMap.get('id');

    if (this.productId) {
      this.isEditMode = true;
      this.loadProduct();
    }

    this.route.queryParams.subscribe((params) => {
      if (params['id']) {
        this.isEditMode = true;
        this.productForm.patchValue({
          name: params['name'],
          price: params['price'],
          available: params['available'],
        });
        this.productId = params['id'];
      }
    });
  }

  loadProduct() {
    if (this.productId) {
      this.productService.getProductById(this.productId).subscribe({
        next: (product: Product) => {
          this.productForm.patchValue(product);
        },
        error: (err) => console.error('Erro ao carregar produto:', err),
      });
    }
  }

  onSubmit() {
    if (this.productForm.invalid) return;
    const productData: Product = this.productForm.value;

    if (this.isEditMode && this.productId) {
      console.log(this.isEditMode)
      console.log(this.productId)
      this.productService.updateProduct(this.productId, productData).subscribe({
        next: () => this.router.navigate(['/products']),
        error: (err) => console.error('Erro ao atualizar produto:', err),
      });
    } else {
      this.productService.createProduct(productData).subscribe({
        next: () => this.router.navigate(['/products']),
        error: (err) => console.error('Erro ao adicionar produto:', err),
      });
    }
  }
}
