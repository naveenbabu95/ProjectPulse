import { Route } from '@angular/router';

export const appRoutes: Route[] = [
    {
		path: '',
		pathMatch: 'full',
		redirectTo: 'board',
	},
	{
		path: 'board',
		loadComponent: () =>
			import('../app/board/board.component').then(
				(m) => m.BoardComponent,
			),
	},
    {
		path: 'detail-view',
		loadComponent: () =>
			import('../app/task-details/task-details.component').then(
				(m) => m.TaskDetailsComponent,
			),
	},
    {
		path: 'page-not-found',
		loadComponent: () =>
			import('../app/page-not-found/page-not-found.component').then(
				(m) => m.PageNotFoundComponent,
			),
	},


];
