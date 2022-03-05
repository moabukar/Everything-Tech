import java.util.*;

public class App {
    public static void main(String[] args) throws Exception {
        Scanner input = new Scanner(System.in);
        System.out.println("This applicaion converts from degrees Celsius to degrees Fahrenheit and vice versa");
        System.out.println("Would you like to input Celsius or Fahrenheit?");
        System.out.println("Enter 'C' for Celsius or 'F' for Fahrenheit: ");    //input must be a single capital letter
        char a = input.next().charAt(0);    //scans for an input taking the first cahracter
        if (a == 'C'){
            System.out.println("Plese eneter the temperature in degrees Celsius (as a number without the sign) that you would like to convert to Fahrenheit");
            try{    //try blcok is used to catch errors from the user input
                int C = input.nextInt();
                Double F = (C*1.8) + 32;    //Satandard calculation to convert from Celsius to Fahrenheit
                System.out.println(F.toString()+"°F");
            }
            catch(Exception e){ //this code will run if an error is encountered in the previous try block
                System.out.println("Please make sure you enter a number");
            }
        }

        else if(a == 'F'){
            System.out.println("Plese eneter the temperature in degrees Fahrenheit (as a number without the sign) that you would like to convert to Celsius");
            try{    //try block is used to catch errors from the user input
                int F = input.nextInt();
                Double C = (F-32) / 1.8;    //Standard calcuation to convert from Fahrenheit to Celsius
                System.out.println(C.toString()+"°C");
            }
            catch(Exception e){ //this code will run if an error is encountered in the previous try block
                System.out.println("Please make sure you enter a number");
            }
        }

        else{
            System.out.println("Sorry input was incorrect application terminating");
        }
        input.close();
    }
}
