 class Person {
        String name;
        int age;
        String address;
        final String country = "indonesia";

        Person(String name, int age, String address) {
            this.name = name;
            this.age = age;
            this.address = address;
        }

        Person(String name, int age) {
            this(name,  age, null);
        }

        Person(String forName) {
            this(null, 0);
        }

        void hello(String paramName){
            System.out.println("Hello " + paramName + " myName " + name);
        }

    } 
