import { Route } from '@angular/router';
import { authGuard } from './core/auth';

export const appRoutes: Route[] = [
  {
    path: 'signup',
    loadComponent: () => import('./signup/signup').then((m) => m.Signup),
  },
  {
    path: '',
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
        canActivate: [authGuard],
        loadChildren: () =>
          import('./profile/profile.routes').then((m) => m.profileRoutes),
      },
      {
        path: 'trips',
        canActivate: [authGuard],
        loadComponent: () =>
          import('./trips/trips.component').then((m) => m.TripsComponent),
      },
      {
        path: 'trips/:id',
        canActivate: [authGuard],
        loadComponent: () =>
          import('./trips/trip-detail/trip-detail.component').then(
            (m) => m.TripDetailComponent,
          ),
      },
    ],
  },
];
