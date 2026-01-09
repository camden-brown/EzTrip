import {
  Component,
  input,
  output,
  ChangeDetectionStrategy,
  computed,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { DateTime } from 'luxon';
import { Activity } from '../../../core/models/trip.model';
import { CategoryIconPipe } from '../../../core/pipes/category-icon.pipe';
import { CategoryLabelPipe } from '../../../core/pipes/category-label.pipe';

@Component({
  selector: 'eztrip-activity-detail-panel',
  standalone: true,
  imports: [
    CommonModule,
    MatIconModule,
    MatButtonModule,
    CategoryIconPipe,
    CategoryLabelPipe,
  ],
  templateUrl: './activity-detail-panel.component.html',
  styleUrl: './activity-detail-panel.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class ActivityDetailPanelComponent {
  activity = input.required<Activity>();
  isOpen = input<boolean>(false);

  close = output<void>();

  formattedTime = computed(() => {
    const dt = DateTime.fromISO(this.activity().time);
    return dt.toFormat('h:mm a');
  });

  formattedDate = computed(() => {
    const dt = DateTime.fromISO(this.activity().time);
    return dt.toFormat('EEEE, MMMM d, yyyy');
  });

  encodeURIComponent(str: string): string {
    return encodeURIComponent(str);
  }

  onBackdropClick(): void {
    this.close.emit();
  }

  onPanelClick(event: Event): void {
    event.stopPropagation();
  }
}
