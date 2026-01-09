import {
  Component,
  input,
  output,
  computed,
  ChangeDetectionStrategy,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { DateTime } from 'luxon';
import { Activity } from '../../../core/models/trip.model';
import { CategoryIconPipe } from '../../../core/pipes/category-icon.pipe';

@Component({
  selector: 'eztrip-activity-card',
  standalone: true,
  imports: [CommonModule, MatIconModule, MatButtonModule, CategoryIconPipe],
  templateUrl: './activity-card.component.html',
  styleUrl: './activity-card.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ActivityCardComponent {
  activity = input.required<Activity>();
  isLast = input<boolean>(false);

  edit = output<Activity>();
  delete = output<Activity>();
  click = output<Activity>();

  formattedTime = computed(() => {
    const dt = DateTime.fromISO(this.activity().time);
    return dt.toFormat('h:mm a');
  });

  getCategoryClass(category: string): string {
    return `category-${category}`;
  }
}
