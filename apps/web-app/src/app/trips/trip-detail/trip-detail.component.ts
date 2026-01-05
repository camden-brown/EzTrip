import { Component, OnInit, signal, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ActivatedRoute, Router } from '@angular/router';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { Trip } from '../../models/trip.model';
import { TripMockService } from '../trip-mock.service';
import { DaySectionComponent } from '../day-section/day-section.component';
import { AiPromptSheetComponent } from '../ai-prompt-sheet/ai-prompt-sheet.component';

@Component({
  selector: 'eztrip-trip-detail',
  standalone: true,
  imports: [
    CommonModule,
    MatIconModule,
    MatButtonModule,
    DaySectionComponent,
    AiPromptSheetComponent,
  ],
  templateUrl: './trip-detail.component.html',
  styleUrl: './trip-detail.component.scss',
})
export class TripDetailComponent implements OnInit {
  private route = inject(ActivatedRoute);
  private router = inject(Router);
  private tripService = inject(TripMockService);

  trip = signal<Trip | null>(null);
  isPromptOpen = signal(false);
  selectedDate = signal<string | null>(null);
  expandedSections = signal<Set<string>>(new Set());
  enableScrollAnimation = signal(true);

  ngOnInit(): void {
    const tripId = this.route.snapshot.paramMap.get('id');
    if (tripId) {
      const trip = this.tripService.getTripById(tripId);
      if (trip) {
        this.trip.set(trip);
      } else {
        this.router.navigate(['/trips']);
      }
    }
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
}
