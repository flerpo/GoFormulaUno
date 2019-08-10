import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { TracksComponent } from './tracks/tracks.component';


import { HttpClientModule } from '@angular/common/http';
import {BrowserAnimationsModule} from '@angular/platform-browser/animations';
import {MatInputModule} from '@angular/material/input';

import {MatButtonModule} from '@angular/material/button';
import {MatCheckboxModule} from '@angular/material/checkbox';
import {MatCardModule} from '@angular/material/card';
import { NavigationComponent } from './navigation/navigation.component';
import { LayoutModule } from '@angular/cdk/layout';
import { MatToolbarModule } from '@angular/material/toolbar';
import { MatSidenavModule } from '@angular/material/sidenav';
import { MatIconModule } from '@angular/material/icon';
import { MatListModule } from '@angular/material/list';
import { MatExpansionModule } from '@angular/material/expansion';
import { PlayersComponent } from './players/players.component';
import { StatisticsComponent } from './statistics/statistics.component';
import { AppRoutingModule } from './app-routing.module';
import { TimeComponent } from './time/time.component';
import { TrackdetailComponent } from './trackdetail/trackdetail.component';
import { FormsModule } from '@angular/forms';


@NgModule({
  declarations: [
    AppComponent,
    TracksComponent,
    NavigationComponent,
    PlayersComponent,
    StatisticsComponent,
    TimeComponent,
    TrackdetailComponent,
    ],
  imports: [
    BrowserModule,FormsModule,
    HttpClientModule,
    BrowserAnimationsModule,
    MatExpansionModule, MatInputModule,
    MatButtonModule, MatCheckboxModule, MatCardModule, LayoutModule, MatToolbarModule, MatSidenavModule, MatIconModule, MatListModule, AppRoutingModule
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
