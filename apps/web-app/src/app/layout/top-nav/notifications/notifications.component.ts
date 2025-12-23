import { Component, Input } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatMenuModule } from '@angular/material/menu';
import { MatIconModule } from '@angular/material/icon';
import { FormsModule } from '@angular/forms';

export interface Notification {
  id: string;
  avatar: string;
  name: string;
  time: string;
  message: string;
  isUnread: boolean;
  hasOnlineStatus?: boolean;
  attachment?: {
    filename: string;
    filesize: string;
  };
}

@Component({
  selector: 'eztrip-notifications',
  standalone: true,
  imports: [CommonModule, MatMenuModule, MatIconModule, FormsModule],
  templateUrl: './notifications.component.html',
  styleUrl: './notifications.component.scss',
})
export class NotificationsComponent {
  @Input() notificationCount = 0;

  activeTab = 'inbox';
  searchQuery = '';

  notifications: Notification[] = [
    {
      id: '1',
      avatar: 'https://i.pravatar.cc/150?img=33',
      name: 'Michael Lee',
      time: '1 hour ago',
      message:
        'You have a new message from the support team regarding your recent inquiry.',
      isUnread: true,
      hasOnlineStatus: true,
      attachment: {
        filename: 'Contract',
        filesize: '2.1 MB',
      },
    },
    {
      id: '2',
      avatar: 'https://i.pravatar.cc/150?img=45',
      name: 'Alice Johnson',
      time: '10 minutes ago',
      message:
        'Your report has been successfully submitted and is under review.',
      isUnread: true,
      hasOnlineStatus: false,
    },
    {
      id: '3',
      avatar: 'https://i.pravatar.cc/150?img=47',
      name: 'Emily Davis',
      time: 'Yesterday at 4:35 PM',
      message:
        'The project deadline has been updated to September 30th. Please check the details.',
      isUnread: true,
      hasOnlineStatus: true,
    },
  ];

  setActiveTab(tab: string): void {
    this.activeTab = tab;
  }

  markAllAsRead(): void {
    this.notifications = this.notifications.map((n) => ({
      ...n,
      isUnread: false,
    }));
  }

  handleNotificationClick(notification: Notification): void {
    // TODO: Implement notification click logic
    console.log('Notification clicked:', notification);
  }

  handleDownload(event: Event, notification: Notification): void {
    event.stopPropagation();
    // TODO: Implement download logic
    console.log('Download clicked:', notification.attachment);
  }
}
