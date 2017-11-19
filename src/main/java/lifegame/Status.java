package lifegame;

import java.util.function.IntPredicate;

public enum Status {

    ALIVE(a -> a == 2 || a == 3),
    NOTHING(a -> a == 3);

    private final IntPredicate predicate;

    private Status(final IntPredicate predicate) {
        this.predicate = predicate;
    }

    public Status nextStatus(final int aliveCount) {
        return predicate.test(aliveCount) ? ALIVE : NOTHING;
    }
}
