import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { TrackdetailComponent } from './trackdetail.component';

describe('TrackdetailComponent', () => {
  let component: TrackdetailComponent;
  let fixture: ComponentFixture<TrackdetailComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ TrackdetailComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(TrackdetailComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
