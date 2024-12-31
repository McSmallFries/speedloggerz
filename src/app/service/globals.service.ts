import {Injectable} from '@angular/core';
import {HttpClient} from '@angular/common/http';
import {User} from '../models/login';
import { SpeedloggerzService } from './speedloggerz.service';



@Injectable({
  providedIn: 'root'
})
export class AppGlobalService {

  isUserLoggedIn: boolean;

  constructor(private service: SpeedloggerzService) {
    
  }

  setIsUserLoggedIn(val: boolean)  {
    this.isUserLoggedIn = val;
  }
}
