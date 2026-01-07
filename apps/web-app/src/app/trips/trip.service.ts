import { inject, Injectable, signal } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';
import { GraphqlService } from '../core/graphql/graphql.service';
import { GET_TRIPS, GET_TRIP, GET_ACTIVITY } from './trip.queries';
import { Trip, Activity } from '../models/trip.model';

interface GetTripsResponse {
  trips: Trip[];
}

interface GetTripResponse {
  trip: Trip;
}

interface GetActivityResponse {
  activity: Activity;
}

@Injectable({
  providedIn: 'root',
})
export class TripService {
  private readonly graphql = inject(GraphqlService);

  private readonly _trips = signal<Trip[]>([]);
  private readonly _isLoadingTrips = signal(false);
  private readonly _tripsError = signal<string | null>(null);

  private readonly _currentTrip = signal<Trip | null>(null);
  private readonly _isLoadingTrip = signal(false);
  private readonly _tripError = signal<string | null>(null);

  readonly trips = this._trips.asReadonly();
  readonly isLoadingTrips = this._isLoadingTrips.asReadonly();
  readonly tripsError = this._tripsError.asReadonly();

  readonly currentTrip = this._currentTrip.asReadonly();
  readonly isLoadingTrip = this._isLoadingTrip.asReadonly();
  readonly tripError = this._tripError.asReadonly();

  loadTrips(): void {
    this._isLoadingTrips.set(true);
    this._tripsError.set(null);

    this.graphql
      .query<GetTripsResponse, void>(GET_TRIPS)
      .pipe(map((response) => response.trips))
      .subscribe({
        next: (trips) => {
          this._trips.set(trips);
          this._isLoadingTrips.set(false);
        },
        error: (err) => {
          console.error('Error loading trips:', err);
          this._tripsError.set('Failed to load trips. Please try again.');
          this._isLoadingTrips.set(false);
        },
      });
  }

  loadTripById(id: string): void {
    this._isLoadingTrip.set(true);
    this._tripError.set(null);

    this.graphql
      .query<GetTripResponse, { id: string }>(GET_TRIP, { id })
      .pipe(map((response) => response.trip))
      .subscribe({
        next: (trip) => {
          this._currentTrip.set(trip);
          this._isLoadingTrip.set(false);
        },
        error: (err) => {
          console.error('Error loading trip:', err);
          this._tripError.set('Failed to load trip. Please try again.');
          this._isLoadingTrip.set(false);
        },
      });
  }

  getTripById(id: string): Observable<Trip> {
    return this.graphql
      .query<GetTripResponse, { id: string }>(GET_TRIP, { id })
      .pipe(map((response) => response.trip));
  }

  getActivityById(id: string): Observable<Activity> {
    return this.graphql
      .query<GetActivityResponse, { id: string }>(GET_ACTIVITY, { id })
      .pipe(map((response) => response.activity));
  }
}
