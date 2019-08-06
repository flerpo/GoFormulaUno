import { Component, OnInit } from '@angular/core';
import { ApiService } from "../api.service";

@Component({
  selector: 'app-players',
  templateUrl: './players.component.html',
  styleUrls: ['./players.component.css']
})
export class PlayersComponent implements OnInit {

  constructor(private apiService: ApiService) { }
  players;
  ngOnInit() {
    this.apiService.getPlayers().subscribe(data => {
      console.log(data)
      this.players = data;
    })
  }

}
