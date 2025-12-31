import { Injectable, inject } from '@angular/core';
import { AuthService as Auth0Service } from '@auth0/auth0-angular';
import { Observable } from 'rxjs';

/**
 * Wraps the Auth0 SDK for easier testing and additional functionality
 */
@Injectable({
  providedIn: 'root',
})
export class AuthService {
  private readonly auth0 = inject(Auth0Service);

  readonly isAuthenticated$ = this.auth0.isAuthenticated$;
  readonly isLoading$ = this.auth0.isLoading$;
  readonly user$ = this.auth0.user$;
  readonly error$ = this.auth0.error$;

  loginWithRedirect(options?: {
    appState?: { target?: string };
  }): Observable<void> {
    return this.auth0.loginWithRedirect(options);
  }

  logout(options?: { logoutParams?: { returnTo?: string } }): Observable<void> {
    return this.auth0.logout(options);
  }

  getAccessTokenSilently(options?: {
    authorizationParams?: {
      audience?: string;
      scope?: string;
    };
  }): Observable<string> {
    return this.auth0.getAccessTokenSilently(options);
  }

  handleRedirectCallback(): Observable<any> {
    return this.auth0.handleRedirectCallback();
  }
}
