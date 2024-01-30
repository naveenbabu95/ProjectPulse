import {Component} from '@angular/core';
import {RouterModule} from '@angular/router';
import {FooterComponent} from './footer/footer.component';
import {HeaderComponent} from './header/header.component';

@Component({
  standalone: true,
  imports: [HeaderComponent, FooterComponent, RouterModule],
  selector: 'project-pulse-root',
  templateUrl: './app.component.html',
  styleUrl: './app.component.scss',
})
export class AppComponent {}
