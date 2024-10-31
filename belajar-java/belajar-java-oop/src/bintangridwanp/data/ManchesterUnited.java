package bintangridwanp.data;

public interface ManchesterUnited extends ManagerMu{

    void Defends();

    void Gelandang();

    void Striker();

    String Kiper();

    void Pemilik();

    default boolean Bermain() {
        return false;
    }
}
