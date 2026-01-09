import { Injectable, inject, signal, computed } from '@angular/core';
import { GraphqlService } from '../graphql/graphql.service';
import { AuthService } from '../auth/auth.service';
import {
  CURRENT_USER_QUERY,
  CREATE_USER_MUTATION,
  type CurrentUserQuery,
  type CreateUserMutation,
} from '../graphql/queries/user.queries';
import type { User, SignupCredentials } from '../models/user.model';
import { Observable, of } from 'rxjs';
import { tap, take, switchMap, map, catchError } from 'rxjs/operators';

/**
 * Service for managing user state and user-related API operations
 */
@Injectable({
  providedIn: 'root',
})
export class UserService {
  private readonly _graphql = inject(GraphqlService);
  private readonly _auth = inject(AuthService);

  private readonly _currentUserSignal = signal<User | null>(null);
  private readonly _loadingSignal = signal<boolean>(false);

  readonly currentUser = this._currentUserSignal.asReadonly();
  readonly loading = this._loadingSignal.asReadonly();
  readonly isUserLoaded = computed(() => this._currentUserSignal() !== null);

  /**
   * Fetches the current user and updates the signal.
   * Automatically checks authentication state before making the API call.
   */
  fetchCurrentUser(): Observable<void> {
    return this._auth.isAuthenticated$.pipe(
      take(1),
      switchMap((isAuth) => {
        if (!isAuth) {
          this._currentUserSignal.set(null);
          return of(void 0);
        }

        this._loadingSignal.set(true);

        return this._graphql
          .query<CurrentUserQuery, void>(CURRENT_USER_QUERY)
          .pipe(
            tap({
              next: (res) => {
                this._currentUserSignal.set(res.currentUser);
                this._loadingSignal.set(false);
              },
              error: (err) => {
                console.error('Failed to fetch current user:', err);
                this._loadingSignal.set(false);
              },
            }),
            catchError(() => of(void 0)),
            map(() => void 0),
          );
      }),
    );
  }

  /**
   * Creates a new user account
   */
  signup(credentials: SignupCredentials): Observable<void> {
    this._loadingSignal.set(true);

    return this._graphql
      .mutate<
        CreateUserMutation,
        { input: SignupCredentials }
      >(CREATE_USER_MUTATION, { input: credentials })
      .pipe(
        tap({
          next: (result) => {
            this._loadingSignal.set(false);
            this._currentUserSignal.set(result.createUser);
          },
          error: () => {
            this._loadingSignal.set(false);
          },
        }),
        map(() => void 0),
      );
  }

  /**
   * Clears the current user state
   */
  clearUser(): void {
    this._currentUserSignal.set(null);
  }
}
