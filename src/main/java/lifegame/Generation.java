package lifegame;

import java.util.Collections;
import java.util.HashMap;
import java.util.Iterator;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.stream.Collectors;

public class Generation {

    private final Field field;
    private final Map<Point, Status> mapping;

    public Generation(final Field field, final List<Status> statusList) {
        this.field = field;

        final Map<Point, Status> m = new HashMap<>();
        final List<Point> points = field.getAllPoints();
        final Iterator<Point> it1 = points.iterator();
        final Iterator<Status> it2 = statusList.iterator();
        while (it1.hasNext() && it2.hasNext()) {
            final Point point = it1.next();
            final Status status = it2.next();
            m.put(point, status);
        }
        this.mapping = Collections.unmodifiableMap(m);
    }

    public Generation nextGeneration() {
        final List<Point> points = field.getAllPoints();

        final List<Status> statusList = points.stream().map(point -> {
            final List<Point> around = field.getAroundPoints(point);
            final int aliveCount = (int) around.stream().map(mapping::get)
                    .filter(a -> a == Status.ALIVE)
                    .count();
            final Status center = mapping.get(point);
            return center.nextStatus(aliveCount);
        }).collect(Collectors.toList());

        return new Generation(field, statusList);
    }

    @Override
    public boolean equals(final Object obj) {
        if (this == obj) {
            return true;
        } else if (obj == null) {
            return false;
        } else if (obj.getClass() != getClass()) {
            return false;
        }
        final Generation other = (Generation) obj;
        return field.equals(other.field) && mapping.equals(other.mapping);
    }

    @Override
    public int hashCode() {
        return Objects.hash(field, mapping);
    }
}
