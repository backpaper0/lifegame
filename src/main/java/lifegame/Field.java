package lifegame;

import java.util.List;

public interface Field {

    List<Point> getAroundPoints(final Point center);

    List<Point> getAllPoints();
}
