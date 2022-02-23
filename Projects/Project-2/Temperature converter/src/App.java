import java.util.*;

public class App {
    public static void main(String[] args) throws Exception {
        Scanner input = new Scanner(System.in);
        System.out.println("This applicaion converts from degrees Celsius to degrees Fahrenheit and vice versa");
        System.out.println("Would you like to input Celsius or Fahrenheit?");
        System.out.println("Enter 'C' for Celsius or 'F' for Fahrenheit: ");
        String a = input.next();
        if (a == "C"){
            System.out.println("Plese eneter the temperature in degrees Celsius (as a number without the sign) that you would like to convert to Fahrenheit");
            //TODO add try catch if user inputs non int value
            int C = input.nextInt();
            Double F = (C*1.8) + 32;    //Satandard calculation to convert from Celsius to Fahrenheit

        }
        if(a == "F"){
            System.out.println("Plese eneter the temperature in degrees Fahrenheit (as a number without the sign) that you would like to convert to Celsius");
            //TODO add try catch if user inputs non int value
            int F = input.nextInt();
            Double C = (F-32) / 1.8;    //Standard calcuation to convert from Fahrenheit to Celsius

        }
        else{
            System.out.println("Sorry input was incorrect application terminating");
        }
    }
}
