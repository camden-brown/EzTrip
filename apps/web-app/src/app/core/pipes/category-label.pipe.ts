import { Pipe, PipeTransform } from '@angular/core';
import { ActivityCategory, CategoryLabel } from '../models/trip.model';

@Pipe({
  name: 'categoryLabel',
  standalone: true,
})
export class CategoryLabelPipe implements PipeTransform {
  transform(category: ActivityCategory): string {
    const labels: Record<ActivityCategory, CategoryLabel> = {
      beach: CategoryLabel.Beach,
      hike: CategoryLabel.Hike,
      food: CategoryLabel.Food,
      hotel: CategoryLabel.Hotel,
      activity: CategoryLabel.Activity,
      transport: CategoryLabel.Transport,
      shopping: CategoryLabel.Shopping,
      entertainment: CategoryLabel.Entertainment,
    };
    return labels[category] || CategoryLabel.Default;
  }
}
