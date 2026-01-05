import {
  Component,
  input,
  output,
  signal,
  computed,
  ChangeDetectionStrategy,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { DateTime } from 'luxon';
import { ItineraryDay } from '../../models/trip.model';
import { ActivityCardComponent } from '../activity-card/activity-card.component';

@Component({
  selector: 'eztrip-day-section',
  standalone: true,
  imports: [
    CommonModule,
    MatIconModule,
    MatButtonModule,
    ActivityCardComponent,
  ],
  templateUrl: './day-section.component.html',
  styleUrl: './day-section.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class DaySectionComponent {
  day = input.required<ItineraryDay>();
  dayNumber = input.required<number>();

  addActivity = output<string>();

  isExpanded = signal(true);

  formattedDate = computed(() => {
    const date = DateTime.fromISO(this.day().date);
    const now = DateTime.now();
    const diffDays = Math.floor(date.diff(now, 'days').days);

    let relativeText = '';

    if (diffDays === 0) {
      relativeText = 'Today';
    } else if (diffDays === 1) {
      relativeText = 'Tomorrow';
    } else if (diffDays === -1) {
      relativeText = 'Yesterday';
    } else if (diffDays > 1 && diffDays <= 7) {
      relativeText = `${diffDays} days from now`;
    } else if (diffDays < -1 && diffDays >= -7) {
      relativeText = `${Math.abs(diffDays)} days ago`;
    }

    const fullDate = date.toFormat('EEEE, MMMM d, yyyy');
    return { fullDate, relativeText };
  });

  toggleExpanded(): void {
    this.isExpanded.update((v) => !v);
  }

  onAddActivity(): void {
    this.addActivity.emit(this.day().date);
  }
}
