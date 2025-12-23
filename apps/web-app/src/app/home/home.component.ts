import { Component, inject } from '@angular/core';
import { CommonModule } from '@angular/common';
import { map, shareReplay, tap } from 'rxjs/operators';
import { GraphqlService } from '../core/graphql/graphql.service';
import {
  CURRENT_USER_QUERY,
  type CurrentUserQuery,
} from '../core/graphql/queries/user.queries';
import type { User } from '../models/user.model';

@Component({
  selector: 'eztrip-home',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './home.component.html',
  styleUrl: './home.component.scss',
})
export class HomeComponent {
  private readonly graphql = inject(GraphqlService);

  readonly currentUser$ = this.graphql
    .query<CurrentUserQuery>(CURRENT_USER_QUERY)
    .pipe(
      map((res) => res.currentUser as User | null),
      tap((u) => console.log('currentUser:', u)),
      shareReplay({ bufferSize: 1, refCount: true }),
    );

  constructor() {
    this.currentUser$.subscribe();
  }
}
