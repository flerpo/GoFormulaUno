import { Component, OnInit, Input } from '@angular/core';
import { Track } from '../track';
import { ApiService } from '../api.service';

@Component({
  selector: 'app-trackdetail',
  templateUrl: './trackdetail.component.html',
  styleUrls: ['./trackdetail.component.css']
})
export class TrackdetailComponent implements OnInit {
  @Input() track: Track;
  
  constructor(private apiService: ApiService) { }

  ngOnInit() {
  }

  onSubmit(track) {
    this.apiService.updateTrack(track).subscribe()
  }
}
