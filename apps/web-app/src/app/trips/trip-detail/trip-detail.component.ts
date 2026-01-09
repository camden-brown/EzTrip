import { Component, OnInit, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { DateTime } from 'luxon';
import { Activity } from '../../core/models/trip.model';
import { TripService } from '../trip.service';
import { DaySectionComponent } from './day-section/day-section.component';
import { AiPromptSheetComponent } from './ai-prompt-sheet/ai-prompt-sheet.component';
import { ActivityDetailPanelComponent } from './activity-detail-panel/activity-detail-panel.component';

@Component({
  selector: 'eztrip-trip-detail',
  standalone: true,
  imports: [
    CommonModule,
    MatIconModule,
    MatButtonModule,
    DaySectionComponent,
    AiPromptSheetComponent,
    ActivityDetailPanelComponent,
  ],
  templateUrl: './trip-detail.component.html',
  styleUrl: './trip-detail.component.scss',
})
export class TripDetailComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  protected readonly tripService = inject(TripService);

  isPromptOpen = signal(false);
  selectedDate = signal<string | null>(null);
  expandedSections = signal<Set<string>>(new Set());
  enableScrollAnimation = signal(true);
  selectedActivity = signal<Activity | null>(null);
  isPanelOpen = signal(false);

  ngOnInit(): void {
    const tripId = this.route.snapshot.paramMap.get('id');
    if (tripId) {
      this.tripService.loadTripById(tripId);
    } else {
      this.router.navigate(['/trips']);
    }
  }

  getClosestDayIndex(): number {
    const trip = this.tripService.currentTrip();
    if (!trip || !trip.itinerary.length) return 0;

    const now = DateTime.now().startOf('day');
    let closestIndex = 0;
    let smallestDiff = Number.MAX_SAFE_INTEGER;

    trip.itinerary.forEach((day, index) => {
      const dayDate = DateTime.fromISO(day.date).startOf('day');
      const diff = Math.abs(dayDate.diff(now, 'days').days);

      if (diff < smallestDiff) {
        smallestDiff = diff;
        closestIndex = index;
      }
    });

    return closestIndex;
  }

  onAddActivity(date: string): void {
    this.selectedDate.set(date);
    this.isPromptOpen.set(true);
  }

  onPromptClose(): void {
    this.isPromptOpen.set(false);
    this.selectedDate.set(null);
  }

  onPromptSubmit(prompt: string): void {
    console.log(
      'AI prompt submitted:',
      prompt,
      'for date:',
      this.selectedDate(),
    );
  }

  goBack(): void {
    this.router.navigate(['/trips']);
  }

  onActivityClick(activity: Activity): void {
    this.selectedActivity.set(activity);
    this.isPanelOpen.set(true);
  }

  onPanelClose(): void {
    this.isPanelOpen.set(false);
    setTimeout(() => {
      this.selectedActivity.set(null);
    }, 300);
  }
}
