import { Component, OnInit } from '@angular/core';
import { ApiService } from "../api.service";

@Component({
  selector: 'app-time',
  templateUrl: './time.component.html',
  styleUrls: ['./time.component.css']
})
export class TimeComponent implements OnInit {

  constructor(private apiService: ApiService) { }
  times;
  tracks;
  players;
  ngOnInit() {
    this.apiService.getTimes().subscribe(data => {
      console.log(data)
      this.times = data;
    })
    this.apiService.getTracks().subscribe(data => {
      console.log(data)
      this.tracks = data;
    })
    this.apiService.getPlayers().subscribe(data => {
      console.log(data)
      this.players = data;
    })
  }

}
