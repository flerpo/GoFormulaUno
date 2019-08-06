import { Component, OnInit } from "@angular/core";
import { ApiService } from "../api.service";

@Component({
  selector: "app-tracks",
  templateUrl: "./tracks.component.html",
  styleUrls: ["./tracks.component.css"]
})
export class TracksComponent implements OnInit {
  tracks;
  constructor(private apiService: ApiService) {}

  ngOnInit() {
    this.apiService.getTracks().subscribe(data => {
      console.log(data);
      
      this.tracks = data;
    });
  }
}
