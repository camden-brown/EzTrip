import { Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { TopNavComponent } from './top-nav/top-nav.component';

@Component({
  selector: 'eztrip-layout',
  standalone: true,
  imports: [CommonModule, RouterModule, TopNavComponent],
  templateUrl: './layout.component.html',
  styleUrl: './layout.component.scss',
})
export class LayoutComponent {}
