import { Injectable, inject, signal } from '@angular/core';
import { GraphqlService } from '../graphql/graphql.service';
import {
  TRIP_SUGGESTION_QUERY,
  type TripSuggestionQuery,
  type TripSuggestionVariables,
} from '../graphql/queries/trip.queries';
import { Observable, of } from 'rxjs';
import { tap, map, catchError } from 'rxjs/operators';

/**
 * Service for managing trip-related operations
 */
@Injectable({
  providedIn: 'root',
})
export class TripService {
  private readonly _graphql = inject(GraphqlService);

  private readonly _loadingSignal = signal<boolean>(false);
  private readonly _suggestionSignal = signal<string>('');

  readonly loading = this._loadingSignal.asReadonly();
  readonly suggestion = this._suggestionSignal.asReadonly();

  /**
   * Generates a trip suggestion based on user prompt.
   */
  generateSuggestion(prompt: string): Observable<string> {
    this._loadingSignal.set(true);

    return this._graphql
      .query<
        TripSuggestionQuery,
        TripSuggestionVariables
      >(TRIP_SUGGESTION_QUERY, { prompt })
      .pipe(
        tap({
          next: (res) => {
            this._suggestionSignal.set(res.tripSuggestion);
            this._loadingSignal.set(false);
          },
          error: (err) => {
            console.error('Failed to get trip suggestion:', err);
            this._suggestionSignal.set(
              'Sorry, something went wrong. Please try again.',
            );
            this._loadingSignal.set(false);
          },
        }),
        catchError(() =>
          of({
            tripSuggestion:
              'Sorry, something went wrong. Please try again.',
          }),
        ),
        map((res) => res.tripSuggestion),
      );
  }

  /**
   * Clears the current suggestion
   */
  clearSuggestion(): void {
    this._suggestionSignal.set('');
  }
}
