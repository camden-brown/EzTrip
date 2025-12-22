import { ChangeDetectionStrategy, Component } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatCardModule } from '@angular/material/card';
import { MatTabsModule } from '@angular/material/tabs';
import { LoginForm, LoginCredentials } from './login-form/login-form';
import { SignupForm, SignupCredentials } from './signup-form/signup-form';

@Component({
  selector: 'eztrip-auth',
  imports: [CommonModule, MatCardModule, MatTabsModule, LoginForm, SignupForm],
  templateUrl: './auth.html',
  styleUrl: './auth.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Auth {
  onLogin(credentials: LoginCredentials) {
    console.log('Login:', credentials);
    // TODO: Implement login logic
  }

  onSignup(credentials: SignupCredentials) {
    console.log('Signup:', credentials);
    // TODO: Implement signup logic
  }
}
