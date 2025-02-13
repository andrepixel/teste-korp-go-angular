import { Injectable } from '@angular/core';
import { Observable } from 'rxjs';
import { ProductGateway } from '../../data/datasources/product-gateway.service';
import { Product } from '../models/product.model';

@Injectable({
  providedIn: 'root',
})
export class ProductService {
  constructor(private productGateway: ProductGateway) {}

  getProducts(): Observable<Product[]> {
    return this.productGateway.getProducts();
  }

  getProductById(productId: string): Observable<Product> {
    return this.productGateway.getProductById(productId);
  }

  createProduct(product: Product): Observable<Product> {
    return this.productGateway.createProduct(product);
  }

  updateProduct(productId: string, product: Product): Observable<Product> {
    console.log('Antes de formatar ->', product);
    const formattedProduct2 = Object.assign({}, product);

    let formattedProduct = {
      ID: product.ID,
      Name: product.Name,
      Price: product.Price,
      Available: product.Available,
    };

    console.log('Depois de formatar ->', formattedProduct2);
    console.log('Testando product.ID ->', product.ID);

    return this.productGateway.updateProduct(productId, formattedProduct);
  }


  deleteProduct(productId: string): Observable<void> {
    return this.productGateway.deleteProduct(productId);
  }
}
