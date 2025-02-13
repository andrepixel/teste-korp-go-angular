import { CommonModule } from '@angular/common';
import { Component } from '@angular/core';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatSnackBar, MatSnackBarModule } from '@angular/material/snack-bar';
import { MatTableModule } from '@angular/material/table';
import { Router } from '@angular/router';
import { Product } from '../../../core/models/product.model';
import { ProductService } from '../../../core/services/product.service';

@Component({
  selector: 'app-product-list',
  standalone: true,
  imports: [
    CommonModule,
    MatTableModule,
    MatCardModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatSnackBarModule,
  ],
  templateUrl: './products-list.component.html',
})
export class ProductsListComponent {
  products: Product[] = [];
  displayedColumns: string[] = ['name', 'price', 'available', 'actions'];
  isDeleting: string | null = null;

  constructor(
    private productService: ProductService,
    private router: Router,
    private snackBar: MatSnackBar
  ) {
    this.loadProducts();
  }

  loadProducts() {
    this.productService.getProducts().subscribe({
      next: (data: Product[]) => {
        this.products = data;
        console.log(this.products);
      },
      error: () => {
        this.showMessage('Erro ao carregar produtos!', true);
      },
    });
  }

  editProduct(product: Product) {
    console.log('Editando produto:', product);

    if (product.ID) {
      this.router.navigate(['product-form'], {
        queryParams: {
          id: product.ID,
          name: product.Name,
          price: product.Price,
          available: product.Available,
        },
      });
    } else {
      this.showMessage('Produto inválido para edição!', true);
    }
  }

  deleteProduct(product: Product) {
    if (confirm('Tem certeza que deseja excluir este produto?')) {
      this.isDeleting = product.ID;

      this.productService.deleteProduct(product.ID).subscribe({
        next: () => {
          this.products = this.products.filter((p) => p.ID !== product.ID);
          this.showMessage('Produto excluído com sucesso!');
          this.isDeleting = null;
        },
        error: () => {
          this.showMessage('Erro ao excluir produto!', true);
          this.isDeleting = null;
        },
      });
    }
  }

  private showMessage(message: string, isError = false) {
    this.snackBar.open(message, 'Fechar', {
      duration: 3000,
      panelClass: isError ? 'error-snackbar' : 'success-snackbar',
    });
  }
}
