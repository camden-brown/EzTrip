import { Pipe, PipeTransform } from '@angular/core';
import { ActivityCategory, CategoryIcon } from '../models/trip.model';

@Pipe({
  name: 'categoryIcon',
  standalone: true,
})
export class CategoryIconPipe implements PipeTransform {
  transform(category: ActivityCategory): string {
    const icons: Record<ActivityCategory, CategoryIcon> = {
      beach: CategoryIcon.Beach,
      hike: CategoryIcon.Hike,
      food: CategoryIcon.Food,
      hotel: CategoryIcon.Hotel,
      activity: CategoryIcon.Activity,
      transport: CategoryIcon.Transport,
      shopping: CategoryIcon.Shopping,
      entertainment: CategoryIcon.Entertainment,
    };
    return icons[category] || CategoryIcon.Default;
  }
}
