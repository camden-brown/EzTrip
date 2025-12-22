import { ChangeDetectionStrategy, Component, signal } from '@angular/core';
import { CommonModule } from '@angular/common';
import { MatIconModule } from '@angular/material/icon';
import { LoginForm, LoginCredentials } from './login-form/login-form';
import { SignupForm, SignupCredentials } from './signup-form/signup-form';

@Component({
  selector: 'eztrip-auth',
  imports: [CommonModule, MatIconModule, LoginForm, SignupForm],
  templateUrl: './auth.html',
  styleUrl: './auth.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Auth {
  isLoginView = signal(true);
  currentYear = new Date().getFullYear();

  onLogin(credentials: LoginCredentials) {
    console.log('Login:', credentials);
    // TODO: Implement login logic
  }

  onSignup(credentials: SignupCredentials) {
    console.log('Signup:', credentials);
    // TODO: Implement signup logic
  }

  switchToSignup() {
    this.isLoginView.set(false);
  }

  switchToLogin() {
    this.isLoginView.set(true);
  }
}
