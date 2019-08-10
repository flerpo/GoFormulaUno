import { Injectable } from "@angular/core";
import { HttpClient } from "@angular/common/http";
import { Track } from './track';
import { map } from 'rxjs/operators';

@Injectable({
  providedIn: "root"
})
export class ApiService {
  constructor(private httpClient: HttpClient) {}
  public getTracks() {
    return this.httpClient.get<Track[]>(`http://localhost:8080/api/v1/tracks/`).pipe(
      map(data => data.map(data => new Track().deserialize(data)))
    );

    // return this.httpClient.get(`http://localhost:8080/api/v1/tracks/`);
  }
  public getPlayers() {
    return this.httpClient.get(`http://localhost:8080/api/v1/players/`);
  }
  public getTimes() {
    return this.httpClient.get(`http://localhost:8080/api/v1/time/`);
  }
}