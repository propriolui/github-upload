import { Component, NgModule, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup, Validators } from '@angular/forms';
import { Router } from '@angular/router';

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {
  form:FormGroup;
  minLength = 8;
  hide = true;


  constructor(private fb:FormBuilder,
              //private authService: AuthService,
              private router: Router) {

                this.form = this.fb.group({
                  email: ['', Validators.required, Validators.email],
                  pwd: ['', Validators.required, Validators.minLength(this.minLength)]
                })
              }

  
  getEmailError() {
    if (this.form.hasError('required')) {
      return 'Questo campo è obbligatorio';
    }


    return this.form.hasError('email') ? 'Formato non valido' : '';
  }

  getPwdError() {
    if (this.form.hasError('required')) {
      return 'Questo campo è obbligatorio';
    }

    if (this.form.hasError('minlength')) {
      return 'Lunghezza minima '+this.minLength+' caratteri'
    }

    return
  }

  onSubmit(){
    const val = this.form.value;

    if(val.email && val.pwd) {
      this.authService.login(val.email, val.password)
        .subscribe(
          () => {
            console.log("User is logged in");
            this.router.navigateByUrl('/');
          }
        )
    }
  }

  ngOnInit(): void {
  }

}
