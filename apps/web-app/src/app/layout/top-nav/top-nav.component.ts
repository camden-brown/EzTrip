import { ChangeDetectionStrategy, Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { MatMenuModule } from '@angular/material/menu';
import { MatBadgeModule } from '@angular/material/badge';
import { MatDividerModule } from '@angular/material/divider';
import { RouterModule } from '@angular/router';
import { NotificationsComponent } from './notifications/notifications.component';
import { ProfileMenuComponent } from './profile-menu/profile-menu.component';
import { AuthService } from '../../core/auth/auth.service';

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
    NotificationsComponent,
    ProfileMenuComponent,
  ],
  templateUrl: './top-nav.component.html',
  styleUrl: './top-nav.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class TopNavComponent {
  private readonly auth = inject(AuthService);

  notificationCount = 3;

  onLogout(): void {
    this.auth
      .logout({
        logoutParams: { returnTo: window.location.origin + '/auth' },
      })
      .subscribe();
  }
}
