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
import { Activity, ActivityCategory } from '../../models/trip.model';

@Component({
  selector: 'eztrip-activity-card',
  standalone: true,
  imports: [CommonModule, MatIconModule, MatButtonModule],
  templateUrl: './activity-card.component.html',
  styleUrl: './activity-card.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ActivityCardComponent {
  activity = input.required<Activity>();
  isLast = input<boolean>(false);

  edit = output<Activity>();
  delete = output<Activity>();

  formattedTime = computed(() => {
    const dt = DateTime.fromISO(this.activity().time);
    return dt.toFormat('h:mm a');
  });

  getCategoryIcon(category: ActivityCategory): string {
    const icons: Record<ActivityCategory, string> = {
      beach: 'beach_access',
      hike: 'terrain',
      food: 'restaurant',
      hotel: 'hotel',
      activity: 'local_activity',
      transport: 'flight',
      shopping: 'shopping_bag',
      entertainment: 'celebration',
    };
    return icons[category] || 'place';
  }

  getCategoryClass(category: ActivityCategory): string {
    return `category-${category}`;
  }
}
