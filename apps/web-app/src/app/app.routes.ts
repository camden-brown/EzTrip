import { Route } from '@angular/router';

export const appRoutes: Route[] = [
  {
    path: '',
    redirectTo: 'auth',
    pathMatch: 'full',
  },
  {
    path: 'auth',
    loadComponent: () => import('./auth/auth').then((m) => m.Auth),
  },
  {
    path: 'dashboard',
    loadComponent: () =>
      import('./layout/layout.component').then((m) => m.LayoutComponent),
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./home/home.component').then((m) => m.HomeComponent),
      },
      // Add more dashboard routes here as the app grows
      // { path: 'trips', loadComponent: () => import('./trips/trips.component').then(m => m.TripsComponent) },
      // { path: 'settings', loadComponent: () => import('./settings/settings.component').then(m => m.SettingsComponent) },
    ],
  },
];
