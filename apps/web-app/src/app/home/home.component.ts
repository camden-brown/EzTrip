import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { map, shareReplay, tap } from 'rxjs/operators';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatChipsModule } from '@angular/material/chips';
import { GraphqlService } from '../core/graphql/graphql.service';
import {
  CURRENT_USER_QUERY,
  type CurrentUserQuery,
} from '../core/graphql/queries/user.queries';
import {
  TRIP_SUGGESTION_QUERY,
  type TripSuggestionQuery,
  type TripSuggestionVariables,
} from '../core/graphql/queries/trip.queries';
import type { User } from '../models/user.model';

@Component({
  selector: 'eztrip-home',
  standalone: true,
  imports: [
    CommonModule,
    FormsModule,
    MatIconModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatChipsModule,
  ],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class HomeComponent {
  private readonly graphql = inject(GraphqlService);

  promptText = '';
  aiResponse = '';
  isLoading = false;

  readonly suggestions = [
    { icon: '🏖️', label: 'Beach getaway' },
    { icon: '🏔️', label: 'Mountain adventure' },
    { icon: '🏙️', label: 'City exploration' },
    { icon: '🎢', label: 'Theme park fun' },
  ];

  readonly currentUser$ = this.graphql
    .query<CurrentUserQuery, void>(CURRENT_USER_QUERY)
    .pipe(
      map((res) => res.currentUser as User | null),
      tap((u) => console.log('currentUser:', u)),
      shareReplay({ bufferSize: 1, refCount: true }),
    );

  readonly greeting$ = this.currentUser$.pipe(
    map((user) => {
      const hour = new Date().getHours();
      let timeGreeting = 'Hello';

      if (hour < 12) timeGreeting = 'Good morning';
      else if (hour < 18) timeGreeting = 'Good afternoon';
      else timeGreeting = 'Good evening';

      const name = user?.firstName || 'there';
      return `${timeGreeting}, ${name}`;
    }),
  );

  constructor() {
    this.currentUser$.subscribe();
  }

  useSuggestion(suggestion: { icon: string; label: string }): void {
    this.promptText = `Plan a ${suggestion.label.toLowerCase()} for my family`;
    this.submitPrompt();
  }

  submitPrompt(): void {
    if (!this.promptText.trim()) return;

    this.isLoading = true;

    this.graphql
      .query<TripSuggestionQuery, TripSuggestionVariables>(
        TRIP_SUGGESTION_QUERY,
        {
          prompt: this.promptText,
        },
      )
      .subscribe({
        next: (res) => {
          this.aiResponse = res.tripSuggestion;
          this.isLoading = false;
        },
        error: (err) => {
          console.error('Failed to get trip suggestion:', err);
          this.aiResponse = 'Sorry, something went wrong. Please try again.';
          this.isLoading = false;
        },
      });
  }

  clearResponse(): void {
    this.aiResponse = '';
    this.promptText = '';
  }
}
