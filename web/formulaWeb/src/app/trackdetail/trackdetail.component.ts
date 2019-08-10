import { Component, OnInit, Input } from '@angular/core';
import { Track } from '../track';

@Component({
  selector: 'app-trackdetail',
  templateUrl: './trackdetail.component.html',
  styleUrls: ['./trackdetail.component.css']
})
export class TrackdetailComponent implements OnInit {
  @Input() track: Track;

  constructor() { }

  ngOnInit() {
    console.log(this.track)
  }

}
