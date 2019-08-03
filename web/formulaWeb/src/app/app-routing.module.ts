import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule, Routes } from '@angular/router';
import { TracksComponent } from './tracks/tracks.component';
import { PlayersComponent } from './players/players.component';
import { StatisticsComponent } from './statistics/statistics.component';
const routes: Routes = [
  { path: 'tracks', component: TracksComponent},
  { path: 'players', component: PlayersComponent },
  { path: 'statistics', component: StatisticsComponent }
];


@NgModule({
  declarations: [],
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule]
})
export class AppRoutingModule { }
