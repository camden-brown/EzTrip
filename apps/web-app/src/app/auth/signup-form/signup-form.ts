import {
  ChangeDetectionStrategy,
  Component,
  output,
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

export interface SignupCredentials {
  firstName: string;
  lastName: string;
  email: string;
  password: string;
}

@Component({
  selector: 'eztrip-signup-form',
  imports: [
    CommonModule,
    ReactiveFormsModule,
    MatFormFieldModule,
    MatInputModule,
    MatButtonModule,
    MatIconModule,
  ],
  templateUrl: './signup-form.html',
  styleUrl: './signup-form.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class SignupForm {
  signup = output<SignupCredentials>();

  hidePassword = signal(true);
  hideConfirmPassword = signal(true);

  signupForm = new FormGroup({
    firstName: new FormControl('', [Validators.required]),
    lastName: new FormControl('', [Validators.required]),
    email: new FormControl('', [Validators.required, Validators.email]),
    password: new FormControl('', [
      Validators.required,
      Validators.minLength(8),
    ]),
    confirmPassword: new FormControl('', [Validators.required]),
  });

  onSubmit() {
    if (this.signupForm.valid) {
      const { password, confirmPassword, ...rest } = this.signupForm.value;
      if (password !== confirmPassword) {
        // TODO: Show error message
        return;
      }
      this.signup.emit({
        ...rest,
        password,
      } as SignupCredentials);
    }
  }

  togglePasswordVisibility() {
    this.hidePassword.update((value) => !value);
  }

  toggleConfirmPasswordVisibility() {
    this.hideConfirmPassword.update((value) => !value);
  }
}
