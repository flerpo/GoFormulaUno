import { Component, OnInit } from "@angular/core";
import { ApiService } from "../api.service";
import { Track } from "../track";

@Component({
  selector: "app-tracks",
  templateUrl: "./tracks.component.html",
  styleUrls: ["./tracks.component.css"]
})
export class TracksComponent implements OnInit {
  selectedTrack: Track;
  public tracks: Track[];
  constructor(private apiService: ApiService) {}

  onSelect(track: Track): void {
    this.selectedTrack = track;
  }
  ngOnInit() {
    this.apiService.getTracks().subscribe(tracks => {
      this.tracks = tracks;
    });
  }
}
