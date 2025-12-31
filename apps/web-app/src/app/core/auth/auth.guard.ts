import { inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { map, take } from 'rxjs';

/**
 * Route guard that requires authentication.
 * Redirects to login if the user is not authenticated.
 */
export const authGuard = () => {
  const auth = inject(AuthService);

  return auth.isAuthenticated$.pipe(
    take(1),
    map((isAuthenticated) => {
      if (!isAuthenticated) {
        auth.loginWithRedirect({
          appState: { target: window.location.pathname },
        });
        return false;
      }
      return true;
    }),
  );
};
