import { Component, NgModule, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  minLength = 8;
  hide = true;
  email = new FormControl('', [Validators.required, Validators.email]);
  pwd = new FormControl('', [Validators.required, Validators.minLength(this.minLength)]);

  form = new FormGroup({
    email: this.email,
    pwd: this.pwd
  })
 

  constructor(
              //private authService: AuthService,
              private router: Router) {}

  ngOnInit(): void{
  }            

  getEmailError(){
    if (this.form.controls["email"].hasError("required")){
      return "Campo richiesto";
    }

    if (this.form.controls["email"].hasError("email")){
      return "Formato non valido";
    }
    return
  }

  getPwdError(){
    if (this.form.controls["pwd"].hasError("required")){
      return "Campo richiesto";
    }
    if (this.form.controls["pwd"].hasError("minlength")){
      return "Lunghezza minima "+ this.minLength + " caratteri";
    }
    return
  }


  onSubmit(){
    const val = this.form.value;

    // if(val.email && val.pwd) {
    //   this.authService.login(val.email, val.password)
    //     .subscribe(
    //       () => {
    //         console.log("User is logged in");
    //         this.router.navigateByUrl('/');
    //       }
    //     )
    // }
  }

}
