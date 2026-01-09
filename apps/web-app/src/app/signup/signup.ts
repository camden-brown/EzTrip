import {
  ChangeDetectionStrategy,
  Component,
  inject,
  OnInit,
  signal,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import {
  FormControl,
  FormGroup,
  ReactiveFormsModule,
  Validators,
} from '@angular/forms';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatCardModule } from '@angular/material/card';
import { Router, RouterModule } from '@angular/router';
import { AuthService } from '../core/auth/auth.service';
import { UserService } from '../core/services/user.service';
import type { SignupCredentials } from '../core/models/user.model';
import type { GraphQLRequestError } from '../core/graphql/graphql-errors';

const PASSWORD_MIN_LENGTH = 8;

@Component({
  selector: 'eztrip-signup',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    RouterModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
    MatCardModule,
  ],
  templateUrl: './signup.html',
  styleUrl: './signup.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class Signup implements OnInit {
  private readonly auth = inject(AuthService);
  private readonly router = inject(Router);
  private readonly userService = inject(UserService);

  currentYear = new Date().getFullYear();
  hidePassword = signal(true);
  hideConfirmPassword = signal(true);
  isSubmitting = signal(false);

  signupForm = new FormGroup({
    firstName: new FormControl('', [Validators.required]),
    lastName: new FormControl('', [Validators.required]),
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [
      Validators.required,
      Validators.minLength(PASSWORD_MIN_LENGTH),
    ]),
    confirmPassword: new FormControl('', [Validators.required]),
  });

  ngOnInit() {
    this.auth.isAuthenticated$.subscribe((isAuthenticated) => {
      if (isAuthenticated) {
        this.router.navigate(['/']);
      }
    });
  }

  onSubmit() {
    const credentials = this.extractFormCredentials();

    this.submitSignup(credentials);
  }

  private extractFormCredentials(): SignupCredentials {
    const { firstName, lastName, email, password } = this.signupForm.value;
    return {
      firstName: firstName!,
      lastName: lastName!,
      email: email!,
      password: password!,
    };
  }

  private submitSignup(credentials: SignupCredentials): void {
    this.isSubmitting.set(true);

    this.userService.signup(credentials).subscribe({
      next: () => this.handleSignupSuccess(),
      error: (err) => this.handleSignupError(err),
    });
  }

  private handleSignupSuccess(): void {
    this.login();
  }

  private handleSignupError(error: GraphQLRequestError): void {
    this.isSubmitting.set(false);
  }

  login() {
    this.auth.loginWithRedirect().subscribe();
  }

  togglePasswordVisibility() {
    this.hidePassword.update((value) => !value);
  }

  toggleConfirmPasswordVisibility() {
    this.hideConfirmPassword.update((value) => !value);
  }
}
