import {CommonModule} from '@angular/common';
import {Component} from '@angular/core';

@Component({
  selector: 'project-pulse-board',
  standalone: true,
  imports: [CommonModule],
  templateUrl: './board.component.html',
  styleUrl: './board.component.scss',
})
export class BoardComponent {}
