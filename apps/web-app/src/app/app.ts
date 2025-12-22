import { Component } from '@angular/core';
import { RouterModule } from '@angular/router';

@Component({
  imports: [RouterModule],
  selector: 'eztrip-root',
  templateUrl: './app.html',
})
export class App {
  protected title = 'web-app';
}
