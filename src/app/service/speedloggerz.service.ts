import { Injectable } from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {User} from '../models/login';
// import { ILogin } from './login';

@Injectable({
  providedIn: 'root'
})
export class SpeedloggerzService {

  databaseURL = `http://127.0.0.1:555`; // server endpoint

  constructor(private http: HttpClient) { }

  async RegisterUser(user: User)  {
    debugger;
    const url = `${this.databaseURL}/register/`;
    await this.http.post(url, user).toPromise().then(p => {
      return p;
    });

  }
}
