import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MatIconModule } from '@angular/material/icon';
import { MatTooltipModule } from '@angular/material/tooltip';

export interface NavItem {
  label: string;
  icon: string;
  route: string;
  badge?: number;
}

export interface NavSection {
  label: string;
  items: NavItem[];
}

@Component({
  selector: 'eztrip-side-nav',
  standalone: true,
  imports: [CommonModule, RouterModule, MatIconModule, MatTooltipModule],
  templateUrl: './side-nav.component.html',
  styleUrl: './side-nav.component.scss',
})
export class SideNavComponent {
  @Input() collapsed = false;
  @Output() navItemClicked = new EventEmitter<void>();

  navSections: NavSection[] = [
    {
      label: 'Home',
      items: [{ label: 'Dashboard', icon: 'dashboard', route: '/dashboard' }],
    },
    // Add more sections as the app grows
    // {
    //   label: 'Planning',
    //   items: [
    //     { label: 'My Trips', icon: 'flight', route: '/dashboard/trips' },
    //     { label: 'Explore', icon: 'explore', route: '/dashboard/explore' },
    //     { label: 'Bookings', icon: 'book_online', route: '/dashboard/bookings' },
    //   ],
    // },
  ];

  onNavItemClick(): void {
    this.navItemClicked.emit();
  }
}
