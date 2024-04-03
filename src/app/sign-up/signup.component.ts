import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import {ILogin, User} from '../models/login';
import { AuthService } from '../service/auth.service';
import {SpeedloggerzService} from '../service/speedloggerz.service';
@Component({
  selector: 'app-login',
  templateUrl: './signup.component.html',
  styleUrls: ['./signup.component.css']
})
export class SignupComponent implements OnInit {

  model: ILogin = { username: 'admin', password: 'admin123' };
  signupForm: FormGroup;
  message: string;
  returnUrl: string;


  constructor(private formBuilder: FormBuilder, private router: Router, public authService: AuthService, private appservice: SpeedloggerzService) { }

  ngOnInit() {
    this.signupForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
    this.returnUrl = '/dashboard';
    this.authService.logout();
  }

  // convenience getter for easy access to form fields
  get getFormControls() { return this.signupForm.controls; }

   async register() {
    if (this.signupForm.invalid) {
        return;
    } else {
    console.log(this.model.username);
    console.log(this.model.password);
    if (this.getFormControls.username.value !== '' && this.getFormControls.password.value !== '') {
      // console.log('Login successful');
      // this.authService.authLogin(this.model);
      // this.
      // localStorage.setItem('isLoggedIn', 'true');
      // localStorage.setItem('token', this.getFormControls.username.value);
      // this.router.navigate([this.returnUrl]);
      const username = this.getFormControls.username.value;
      const password = this.getFormControls.password.value;
      const usr = new User();
      usr.Username = username;
      usr.Password = password;
      await this.appservice.RegisterUser(usr);

    } else {
      this.message = 'Please check your username and password';
    }
  }
}
}
