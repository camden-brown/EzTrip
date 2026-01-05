import { Route } from '@angular/router';

export const profileRoutes: Route[] = [
  {
    path: '',
    loadComponent: () =>
      import('./profile.component').then((m) => m.ProfileComponent),
    children: [
      {
        path: 'billing',
        loadComponent: () =>
          import('./billing/billing.component').then((m) => m.BillingComponent),
      },
    ],
  },
];
