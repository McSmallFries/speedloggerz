import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';
import { ILogin } from '../models/login';
import { AuthService } from '../service/auth.service';
@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  model: ILogin = { username: 'admin', password: 'admin123' };
  loginForm: FormGroup;
  message: string;
  returnUrl: string;


  constructor(private formBuilder: FormBuilder, private router: Router, public authService: AuthService) { }

  ngOnInit() {

    if (localStorage.getItem('isLoggedIn') === 'true') {
        this.router.navigate(['/dashboard']);
      } else {
          this.router.navigate(['/login']);
        }

    this.loginForm = this.formBuilder.group({
      username: ['', Validators.required],
      password: ['', Validators.required]
    });
    this.returnUrl = '/dashboard';
    this.authService.logout();
  }

  // convenience getter for easy access to form fields
  get getLoginFormControls() { return this.loginForm.controls; }

   login() {
    if (this.loginForm.invalid) {
        return;
    } else {
    console.log(this.model.username);
    console.log(this.model.password);
    if (this.getLoginFormControls.username.value !== '' && this.getLoginFormControls.password.value !== '') {
      // console.log('Login successful');
      // this.authService.authLogin(this.model);

      localStorage.setItem('isLoggedIn', 'true');
      localStorage.setItem('token', this.getLoginFormControls.username.value);
      this.router.navigate([this.returnUrl]);
    } else {
      this.message = 'Please check your username and password';
    }
  }
}
}
