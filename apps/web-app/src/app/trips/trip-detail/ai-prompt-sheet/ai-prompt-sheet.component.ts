import {
  Component,
  output,
  signal,
  ChangeDetectionStrategy,
  ElementRef,
  viewChild,
  AfterViewInit,
} from '@angular/core';
import { CommonModule } from '@angular/common';
import { FormsModule } from '@angular/forms';
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';

@Component({
  selector: 'eztrip-ai-prompt-sheet',
  standalone: true,
  imports: [CommonModule, FormsModule, MatIconModule, MatButtonModule],
  templateUrl: './ai-prompt-sheet.component.html',
  styleUrl: './ai-prompt-sheet.component.scss',
  changeDetection: ChangeDetectionStrategy.OnPush,
})
export class AiPromptSheetComponent implements AfterViewInit {
  close = output<void>();
  submit = output<string>();

  promptText = signal('');
  isVisible = signal(false);
  textareaRef = viewChild<ElementRef<HTMLTextAreaElement>>('promptInput');

  ngAfterViewInit(): void {
    setTimeout(() => {
      this.isVisible.set(true);
      this.focusInput();
    }, 50);
  }

  focusInput(): void {
    this.textareaRef()?.nativeElement.focus();
  }

  onBackdropClick(): void {
    this.closeSheet();
  }

  closeSheet(): void {
    this.isVisible.set(false);
    setTimeout(() => {
      this.close.emit();
    }, 300);
  }

  onSubmit(): void {
    const text = this.promptText().trim();
    if (text) {
      this.submit.emit(text);
      this.closeSheet();
    }
  }

  onKeydown(event: KeyboardEvent): void {
    if (event.key === 'Enter' && !event.shiftKey) {
      event.preventDefault();
      this.onSubmit();
    }
    if (event.key === 'Escape') {
      this.closeSheet();
    }
  }

  updatePrompt(event: Event): void {
    const target = event.target as HTMLTextAreaElement;
    this.promptText.set(target.value);
  }
}
