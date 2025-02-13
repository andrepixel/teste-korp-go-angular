import { Routes } from '@angular/router';
import { HomeComponent } from './presentation/pages/home/home.component';
import { ProductsListComponent } from './presentation/pages/products-list/products-list.component';

export const routes: Routes = [
  { path: '', redirectTo: 'home', pathMatch: 'full' },
  { path: 'home', component: HomeComponent },
  { path: 'products', component: ProductsListComponent },
  {
    path: 'product-form',
    loadComponent: () =>
      import(
        './presentation/components/product-form/product-form.component'
      ).then((m) => m.ProductFormComponent),
  },
  {
    path: 'product-form/:id',
    loadComponent: () =>
      import(
        './presentation/components/product-form/product-form.component'
      ).then((m) => m.ProductFormComponent),
  },
];
