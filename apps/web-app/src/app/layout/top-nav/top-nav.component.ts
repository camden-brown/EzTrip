import { Component, EventEmitter, Input, Output } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatMenuModule } from '@angular/material/menu';
import { MatBadgeModule } from '@angular/material/badge';
import { MatDividerModule } from '@angular/material/divider';
import { RouterModule } from '@angular/router';

@Component({
  selector: 'eztrip-top-nav',
  standalone: true,
  imports: [
    CommonModule,
    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatMenuModule,
    MatBadgeModule,
    MatDividerModule,
    RouterModule,
  ],
  templateUrl: './top-nav.component.html',
  styleUrl: './top-nav.component.scss',
})
export class TopNavComponent {
  @Input() isMobile = false;
  @Output() menuToggle = new EventEmitter<void>();

  // Mock user data - replace with actual auth service
  user = {
    name: 'John Doe',
    email: 'john.doe@example.com',
    avatar: null as string | null,
  };

  notificationCount = 3;

  onMenuToggle(): void {
    this.menuToggle.emit();
  }

  onLogout(): void {
    // TODO: Implement logout logic
    console.log('Logout clicked');
  }
}
