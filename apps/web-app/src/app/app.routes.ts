import { Route } from '@angular/router';
import { authGuard } from './core/auth';

export const appRoutes: Route[] = [
  {
    path: 'auth',
    loadComponent: () => import('./auth/auth').then((m) => m.Auth),
  },
  {
    path: '',
    canActivate: [authGuard],
    loadComponent: () =>
      import('./layout/layout.component').then((m) => m.LayoutComponent),
    children: [
      {
        path: '',
        loadComponent: () =>
          import('./home/home.component').then((m) => m.HomeComponent),
      },
      {
        path: 'profile',
        loadChildren: () =>
          import('./profile/profile.routes').then((m) => m.profileRoutes),
      },
      {
        path: 'trips',
        loadComponent: () =>
          import('./trips/trips.component').then((m) => m.TripsComponent),
      },
      {
        path: 'trips/:id',
        loadComponent: () =>
          import('./trips/trip-detail/trip-detail.component').then(
            (m) => m.TripDetailComponent
          ),
      },
    ],
  },
];
