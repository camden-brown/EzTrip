import { Component, inject, ChangeDetectionStrategy } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { map, take, switchMap } from 'rxjs/operators';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { MatChipsModule } from '@angular/material/chips';
import { AuthService } from '../core/auth/auth.service';
import { UserService } from '../core/services/user.service';
import { TripService } from '../core/services/trip.service';

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
  private readonly auth = inject(AuthService);
  private readonly userService = inject(UserService);
  private readonly tripService = inject(TripService);

  promptText = '';

  readonly suggestions = [
    { icon: '🏖️', label: 'Beach getaway' },
    { icon: '🏔️', label: 'Mountain adventure' },
    { icon: '🏙️', label: 'City exploration' },
    { icon: '🎢', label: 'Theme park fun' },
  ];

  readonly isAuthenticated$ = this.auth.isAuthenticated$;
  readonly currentUser = this.userService.currentUser;
  readonly userLoading = this.userService.loading;
  readonly tripSuggestion = this.tripService.suggestion;
  readonly tripLoading = this.tripService.loading;

  readonly greeting$ = this.isAuthenticated$.pipe(
    map((isAuth) => {
      const hour = new Date().getHours();
      let timeGreeting = 'Hello';

      if (hour < 12) timeGreeting = 'Good morning';
      else if (hour < 18) timeGreeting = 'Good afternoon';
      else timeGreeting = 'Good evening';

      if (!isAuth) return timeGreeting;

      const user = this.currentUser();
      const name = user?.firstName || 'there';
      return `${timeGreeting}, ${name}`;
    }),
  );

  constructor() {
    this.userService.fetchCurrentUser().subscribe();
  }

  useSuggestion(suggestion: { icon: string; label: string }): void {
    this.promptText = `Plan a ${suggestion.label.toLowerCase()} for my family`;
    this.submitPrompt();
  }

  submitPrompt(): void {
    if (!this.promptText.trim()) return;

    this.auth.isAuthenticated$
      .pipe(
        take(1),
        switchMap((isAuth) => {
          if (!isAuth) {
            return this.auth.loginWithRedirect({
              appState: { target: '/' },
            });
          }

          return this.tripService.generateSuggestion(this.promptText);
        }),
      )
      .subscribe();
  }

  clearResponse(): void {
    this.tripService.clearSuggestion();
    this.promptText = '';
  }
}
