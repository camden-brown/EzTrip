import { Component, signal, ViewChild } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { MatSidenav, MatSidenavModule } from '@angular/material/sidenav';
import { BreakpointObserver, Breakpoints } from '@angular/cdk/layout';
import { TopNavComponent } from './top-nav/top-nav.component';
import { SideNavComponent } from './side-nav/side-nav.component';

@Component({
  selector: 'eztrip-layout',
  standalone: true,
  imports: [
    CommonModule,
    RouterModule,
    MatSidenavModule,
    TopNavComponent,
    SideNavComponent,
  ],
  templateUrl: './layout.component.html',
  styleUrl: './layout.component.scss',
})
export class LayoutComponent {
  @ViewChild('sidenav') sidenav!: MatSidenav;

  isMobile = signal(false);
  sidenavOpened = signal(true);

  constructor(private breakpointObserver: BreakpointObserver) {
    this.breakpointObserver
      .observe([Breakpoints.XSmall, Breakpoints.Small])
      .subscribe((result) => {
        this.isMobile.set(result.matches);
        if (result.matches) {
          this.sidenavOpened.set(false);
        } else {
          this.sidenavOpened.set(true);
        }
      });
  }

  toggleSidenav(): void {
    if (this.isMobile()) {
      this.sidenav.toggle();
    } else {
      this.sidenavOpened.update((v) => !v);
    }
  }

  closeSidenavOnMobile(): void {
    if (this.isMobile()) {
      this.sidenav.close();
    }
  }
}
